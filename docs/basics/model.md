# Models

Asher provides scaffolding of models with a few bindings to relationships and columns in your database. You can specify 
a list of columns that must be included for CRUD operations, list of columns that must not be included in the response 
of API. Basically scaffolding a model builds the foundation on which you can further develop your application.

### Basics
To scaffold your models, you'll need to specify the details of them in your `config.asher`.
All the models must be listed under `model` key of your root configuration. 

Here's an example of a model configuration:

```json
{
  "models": [{
    "name":"id",
    "primary" : true,
    "fillable": true
  }, {
    "name": "password",
    "hidden": "true"
  }] 
}
```