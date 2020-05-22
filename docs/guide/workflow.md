# The Transactor Pattern
## Why and what is it ?
Our team wrote backends in spring boot before switching over to Laravel. We were fans of the DAO service pattern and 
implemented something very similar in Laravel. 
Laravel offered us an ORM out of the box. So we didnt necessarily have to build the Repository layer. We just wrapped our 
queries in service methods, including create and update. This was called in our controllers. We had several MVC and 
RESTful project, sometimes both. This seemed to be a good approach, with all our queries and persistence login written 
in one central place until the service class became really fat.

We had written a few projects in rails, and we really liked the persistence validation that rails' ActiveRecord provides. 
Laravel performs validations on the controller with FormRequests, which, imo is insufficient, since data could change in
between the controller and the persistence layer. We realized we needed persistence validations, which could be done by 
writing custom validators. Now this had to be generalized for all models and that's how this pattern came to be. 
Also, we got bored of writing essentially the same CUD methods, so we tried to think of a way to automate 
this portion of the API. 

The transactor pattern takes heavy inspiration from GraphQL and CQRS design pattern for distributed systems. 
We essentially wanted to build a CQRS for monoliths. 

## Replacement for existing patterns?
Do note, this is not a replacement for MVC And MVP. Its just a set of guidelines that organizes your code in a cleaner 
fashion by separating the *create*, *update* and *delete* operations. This means this can definitely be used with both 
the aforementioned patterns as we have done several times.

## Components
The Transactor pattern has the following components
- Models
- Mutators
- Transactors
- Queries

**Models** are your database classes that find particular data you're looking for, update that data, or remove data. In 
our approach we try to use the model just as a data mapper. They must have ```getCreateValidationRules()``` and 
```getUpdateValidationRules()``` defined in them as a static method.

**Mutator** Every model must have its own mutator file. The mutator performs all write operations on the database ie Create, 
Update and Delete. Mutators aren't allowed to call other mutators. They perform *persistence validation*. You can persist
data only through the mutator.

**Transactor** Wraps a mutator inside a transaction. It is allowed to call other transactors if 
necessary. 

**Queries** Class that contains all your read queries.
 
Now your controllers perform another validation checking for the presence of keys in the input. 
 
 For visual aid<br>
 ![image](https://drive.google.com/uc?export=view&id=1FCiPFjxNxx-X8P8An_qD1RTIobBDAeS-)
 
### How is this similar to CQRS ?
Since Transactors handle the *Commands*, (Create, Update and Delete) and *Queries* handle the Query, we have technically
the DNA of this pattern. All we need is a event bus and (2) separate databases.
 
## Our Implementation
**This implementation is PHP SPECIFIC**
We are using the model classes that php provides.
The model class has a few things defined.
* a create validation rules array that holds the validation rules required in the create method
```
array(
        'valid_till'    => 'required|date_format:Y-m-d H:i:s',
        'course_name'   => 'required|max:255|unique:courses',
        "price"         => "required|numeric|min:1",
        'file_urls'      => 'required|array',
        'description'   => 'sometimes|required',
        'created_by'    => 'required|exists:users,id',
);
``` 
* an update validation rules an associative array that holds the validation rules required to update and (soft)delete 
the model.
```
array(
        'valid_till'    => 'required|date_format:Y-m-d H:i:s',
        'course_name'   => 'required|max:255|unique:courses,course_name,'.ids['courses'],
        "price"         => "required|numeric|min:1",
        'file_urls'      => 'required|array',
        'description'   => 'sometimes|required',
        'updated_by'    => 'exists:users,id',
        'deleted_at'    => 'required|date_format:Y-m-d H:i:s',
    );
```
In our implementation, keys present in the update validation rules array are *optional*, ie we don't require all of them
when the method is called. We merge the values present in the updateValidationRules array with the input array and 
persist the data. Our controller decides which keys are required for a particular endpoint, and hence it has its own 
form request validator. If you notice we appends `ids['courses]` to the unique validation rule to ignore the current row
id. For more details on laravel validations click [here](https://laravel.com/docs/7.x/validation#rule-unique).
 
Keys present in the create validation rules array are *compulsory*.

##### Mutators
We have a BaseMutator class that takes the fully qualified name of the model as input and defines create, update and 
delete methods. All models extend the BaseMutator and pass its `fullyQualifiedName` to the constructor.
The BaseMutator provides basic create, update and (soft)delete methods and runs validations before persisting anything.

##### Transactors- 
We have a BaseTransactor, that takes a `BaseQuery` and a `BaseMutator` reference as input. `BaseMutator` provides
the transactor with the methods necessary to persist data. 
Every model has its own transactor, with a reference to its mutator. These are injected as singletons.
It is here where you write most of your business logic. Transactors can call other transactors and hence inject their 
references in the constructor of another dependent transactor.

We also have a variant of BaseTransactor, the `FileUploadTransactor`. It allows for models to upload file and perform CUD 
operations. It asks for a `BaseUploadHelper` instance in its constructor along with `BaseQuery` and `BaseMutator` and 
delegates CUD to the parent BaseTransactor and file upload operations to the BaseUploadHelper. It provides you with 
methods to create, update and delete files along with models.

##### Query
We have a `BaseQuery` class that takes the `fullyQualifiedName` of the model as input. It is here where we write most 
of our read queries. We haven't given this layer much thought at the time of writing this. It provides `paginate()`,
`fetchOneByCol` and `fetchOneByIdWith` methods as of now.
 

