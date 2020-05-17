# The Transactor Pattern
## Why and what is it ?
The transactor pattern takes heavy inspiration from the CQRS design pattern for distributed systems. We essentially wanted
to build that for monoliths. We also liked the persistence validation that rails' ActiveRecord provides. Laravel performs
validations on the controller with FormRequests, which is, imo an anti-pattern. We wanted to perform persistence validations,
which could be done by writing custom validators, but wanted to generalize this for all models and that's how this pattern
came to be. Also, we got bored with writing essentially the same CUD methods, so we tried to think of a way to automate 
this portion of the API. 

## Replacement for existing patterns?
Do note, this is not a replacement for MVC And MVP. It just organizes your code in a cleaner fashion by separating the 
*create*, *update* and *delete* operations. This means this can definitely be used with both the aforementioned patterns
as we have done several times.

## Components
The Transactor pattern has the following components
- Models
- Mutators
- Transactors
- Queries

**Models** are your database classes that find particular data you're looking for, update that data, or remove data. In 
our approach we try to use the model just as a data mapper. They must have ```getCreateValidationRules()``` and 
```getUpdateValidationRules()``` defined in them as a static method.

**Mutator** Every model has its own mutator file. The mutator performs all write operations on the database ie Create, 
Update and Delete. Mutators aren't allowed to call other mutators. They perform *persistence validation*.

**Transactor** Wraps a mutator inside a transaction. Performs logging. It is allowed to call other transactors if 
necessary. 

**Queries** Class that contains all your read queries.
 
 For visual aid
 ![image](https://drive.google.com/uc?export=view&id=1feXRRJbvQThx8KF9hkrFvNc2wh3nZxHL)
 
###How is this similar to CQRS ?
Since Transactors handle the *Commands*, (Create, Update and Delete) and *Queries* handle the Query, we have technically
the DNA of this pattern. All we need is a event bus and 2 separate databases.

##Our Implementation
**This implementation is PHP SPECIFIC**
We have a BaseMutator class that takes the fully qualified name of the model as input and defines create, update and 
delete methods.
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
* an update validation rules an associative array that holds the validation rules required to update and (soft)delete the model.
```
array(
        'valid_till'    => 'required|date_format:Y-m-d H:i:s',
        'course_name'   => 'required|max:255|unique:courses',
        "price"         => "required|numeric|min:1",
        'file_urls'      => 'required|array',
        'description'   => 'sometimes|required',
        'updated_by'    => 'exists:users,id',
        'deleted_at'    => 'required|date_format:Y-m-d H:i:s',
    );
```
Keys present in the update validation rules array are *optional*, ie we don't require all of them when the method is called.
Keys present in the create validation rules array are *compulsory*.