# fam-flavour-go

## usage

// cat example.yml | go run add.go

## tasks

* the ordering of the yaml is changed on writing but it should stay the same. i should changes the datatypes to a MapSlice https://stackoverflow.com/questions/33639269/preserving-order-of-yaml-maps-using-go how?!
* small changes are done to the content during writing: True -> true. maybe this can also be resolved with the Mapslice? idk!
* the "addon" could maybe not exist in the target yaml. this is currently not handled and will create an error
* the code is horrible, help
