# Controller

### Table of Contents

* [Basic Usage](#basic-usage)
* [HttpMethods](#httpmethods)
* [Type](#type)

### Basic Usage
To scaffold a controller all you need to do is add a `controller` key in the `config.asher` file with some properties. As shown below:
```json
{
  ...
  "controller": {
    "rest": true
  },
  ...
}
```
This is the basic configuration of a RESTful controller.

When you put this configuration the asher generates a controller in  `App\Http\Controllers\Api` with the model name you specified with `5 HTTP methods`:
* create `POST`
* update `PUT`
* delete `DELETE`
* findById `GET`
* getAll  `GET`

With CRUD operations performed in it. So you need to focus on the main business logic.

Every time you add/update a `controller` under the `models` list you need to run following command
```bash
$ asher scaffold
```

### HttpMethods

There are some case where you don't need all the functions inside the controller. We provide you with that flexibility so that your code looks clean. You just need to add a key `httpMethods` inside the `controller` key by default we add all the methods needed. To do so:
```json
{
  "controller": {
    "rest": true,
    "httpMethods" : [
      "POST",
      "GET",
      "PUT",
      "DELETE"
    ]   
  }

}
```

`NOTE: By default all the methods are scaffolded`

After you add/update the key please run the following command:
```bash
$ asher scaffold
```

### Type
 
There may be case where you need to upload files or images we also provide customization according to the demands.You just need to specify `type` key inside the `controller` To do so:

```json
{
  "controller": {
    "type": "file"  
  }

}
```
`Note: By default the type is default`

After you add/update the key please run the following command:
```bash
$ asher scaffold
```

We provide 3 different `types` of transactors:
* [default](#)
* [file](#)
* [image](#)