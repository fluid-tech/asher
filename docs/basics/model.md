# Models and Migrations

Asher provides you with the scaffolding of `Eloquent\Model` and `Migrations` of laravel. It generates all the specified 
models configured with the relationships and the constraints. Hence, saving your time by generating all the boiler plate
code, so you can focus on the business logic.

#### Table of content
* [Basic Usage](/basics/model#configuration) 
* [Example](/basics/model#example) 
* [Soft Deletes](#) 
* [Audit Columns](#) 
* [Time stamps](#) 

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