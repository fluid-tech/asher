# Models and Migrations

Asher provides you with the scaffolding of `Eloquent\Model` and `Migrations` of laravel. It generates all the specified 
models configured with the relationships and the constraints. Hence, saving your time by generating all the boiler plate
code, so you can focus on the business logic.

#### Table of content
* [Basic Usage](/basics/model#configuration) 
* [Example](/basics/model#example)
* [Audit Columns](#)
* [Time stamps](#) 
* [Soft Deletes](#) 
* [Columns](/basics/model#columns)
    * [Validations](/basics/model#validations)
    * [Fillable](#) 
    * [Hidden](#) 
    * [Column Types](#)

### Basic Usage
To scaffold a model, the first thing you need to do is specify the details of your model in `config.asher`
To do this, you need to add `models` key in your configuration file with some properties. As shown below:

```json
{
  "models": [{
    "name": "model_name",
    "cols": [{
      "name": "column_name",
      "colType": "datatype"
    }]  
  }]
}
```
Each model object must contain a `name` that specifies the name for that model and `cols` which is an array of object
that describes the columns required in that model. Each `cols` object must have `name` to describe the name of the column,
`colType` which is the datatype for that column. You can see a list of data types supported by asher [here]()

Every time you add/update a model under the `models` list. You must run the following command to create/update the 
corresponding laravel `Eloquent\Model` and `Migrations`:
```bash
$ asher scaffold
```


### Example
Let us consider that we need to generate a `User` model with the following configuration:
 * Table name - `users`
 * Columns
    * `id` - The primary key for this table of type bigInteger.
    * `user_name` - A string of size 255, that stores the name of the user.
    * `password` - A hidden field that stores the password of the user, with type string and size of 12 characters.

To scaffold this Model & Migration with the help of asher, follow the steps below:
1. First you'll need to add the configuration shown below in your `config.asher` 
```json
{
      "models" : [{
        "name": "users",
        "cols": [{
          "name": "id",
          "colType": "bigInteger",
          "primary": true
        }, {
          "name": "user_name",
          "colType": "string",
          "validations": "max:255",
          "fillable": "true"
        }, {
          "name": "password",
          "colType": "string",
          "validations": "min:8|max:12",
          "hidden": "true"
        }],
        "timestamps": true
      }] 
}
```

2. Run the following command in the base directory of your project:
```bash
$ asher scaffold
```

That's it! Asher will now generate the following Model and Migration files for the specified model.

### Columns
#### Validations
One of the most important part of an application is validation of the data. To make sure that the input data provided by
the user is correct and would not affect the consistency of our database. Asher provides validation of the columns too, 
according to our [Transactor Pattern]() we recommend storing the validations of each column in the model itself. 
This helps asher in generating the basic CRUD methods for your model.

Every `cols` object in a model has 2 keys `createRules` and `updateRules` that accepts a string specifying the rules for
that column. As shown below:
```json
{
  "models": [{
    "name": "model_name",
    "cols": [{
      "name": "column_name",
      "createRules": "validation_rules",
      "updateRules": "validation_rules"
    }]
  }]
}
```
Asher supports all the validation rules provided by laravel. You just need to provide a string of all the 
rules to `createRules` and `updateRules`. You can view all the available rules [here](https://laravel.com/docs/7.x/validation#available-validation-rules).

##### Example
Here's an example configuration column `email` that must be unique string with a max size of 255.
```json
{
  "models": [{
    "name": "users",
    "cols": [{
      "name": "user_email",   
      "colType": "string",
      "createRules": "string|max:255|unique:users",
      "updateRules": "string|max:255|unique:users"
    }]
  }]
}
```  


#### Fillable
To generate the CRUD methods for your model, you must specify asher about the columns that are expected in the input.
You can do this by setting the property `fillable` as `true` under the object of the respective column. Asher adds these
list of columns to the fillable array of that Model, so you can perform create/update operations on that column.

##### Example 
```json
{
  "models": [{
    "name": "users",
    "cols": [{
      "name": "user_email",
      "colType": "string",
      "fillable": "true"
    }]
  }]
}
```

#### Hidden
Sometimes you need to hide the values of some columns from users as they may consider sensitive information 
(E.g: passwords). Asher allows you to configure these columns so that they are hidden in the response. You can do this 
by simply setting the property `hidden` as `true under the object of the respective column.

##### Example
```json
{
  "models": [{
    "name": "users",
    "cols": [{
      "name": "user_email",
      "colType": "string",
      "hidden": "true"
    }]
  }]
}
```

#### Column Types / Default Value
You can specify the type of your column using the property `colType` under each object of the respective column. Each
column must specify its type. The `cols` object has another property `defaultVal` that lets you specify the default 
value of that column.

To specify the **size** of a column, you need to add a createValidation rule `max:size` describing the required size.  

##### Example
Here's an example of column `email` of type string of size 255 with default value of `example@gmail.com` 
```json
{
  "models": [{
    "name": "users",
    "cols": [{
      "name": "user_email",
      "colType": "string",
      "createRules": "max:255",
      "defaultVal": "example@gmail.com"
    }]
  }]
}
```

##### Supported list of columns
Asher currently supports the following data types of columns.


| Type                | Description                                                                                    |
|---------------------|------------------------------------------------------------------------------------------------|
| `string`            | VARCHAR equivalent of column.                                                                  |  
| `integer`           | INTEGER equivalent column.                                                                     |
| `boolean`           | BOOLEAN equivalent column.                                                                     |
| `bigInteger`        | BIGINT equivalent column.                                                                      |
| `date`              | DATE equivalent column.                                                                        |
| `enum`              | ENUM equivalent column. To specify the allowed values in a enum, you can use the `allowed` property under a object of `cols`|