Hello, it is my token-generator,I create this project with golang,gin,sqlite(in beta), later I add postgres.
Okay,how this rest api work?
you send a POST request , for example:
{
"message":"test"
}
and my project create a unique token and save your message with token like a key-value system
if u want a get your message you need the set GET request and write your token in URL
after this you can have your message back)
