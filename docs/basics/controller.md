#Controller
###What is REST API?
Let’s say you’re trying to find videos about Batman on Youtube. You open up Youtube, type "Batman" into a search field, hit enter, and you see a list of videos about Batman. A REST API works in a similar way. You search for something, and you get a list of results back from the service you’re requesting from.


An `API` is an application programming interface. It is a set of rules that allow programs to talk to each other. The developer creates the API on the server and allows the client to talk to it.

`REST` determines how the API looks like. It stands for “Representational State Transfer”. It is a set of rules that developers follow when they create their API. One of these rules states that you should be able to get a piece of data (called a resource) when you link to a specific URL.

Each URL is called a `request` while the data sent back to you is called a `response`.


Asher provides you with the scaffolding of Laravel `RESTful Controller` as of now. 


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

###HttpMethods

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

###Type
 
There may be case where you need to upload files or images we also provide customization according to the demands.You just need to specify `type` key inside the `controller` To do so:

```json
{
  "controller": {
    "type": "file"  
  }

}
```
`Note: By default the type is default`

We provide 3 different types of transactors:
* [default](#)
* [file](#)
* [image](#)