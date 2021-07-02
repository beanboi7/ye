# ye
<ul>
<li>CLI which gives random quotes of Kanye, built using GO.

<li>A CLI has 3 parts apart from the rootword (ye), namely
</ul>

> **Command**
> **Argument** 
> --**Flag** <br>
> Rootword for this CLI - _ye_
- The ye CLI has only 1 command, ```raNdom```, which is self explanatory.

## How to build and use the CLI:
> I assume that you've the GO compiler installed in your system, else no worries just Google about it and get it installed.
- Open your terminal and run the ```go get github.com/beanboi7/ye``` command in a convenient directory.
- Now build the CLI using ```go build main.go``` and you should see a ```main.exe``` in the current directory.
- Type the root word ```ye``` followed by the command ```raNdom``` and a Kanye quote pops out.


### Note:

- In linux distros, before running ```go get github.com/beanboi7/ye```, don't forget to export PATH variable that points to bin folder.
- (ignore this step if PATH is already set)

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
```

#### To contribute to the code/ fix any issues make a pull request.

#### This CLI was made possible thanks to the API made by [Andrew Jazbec](https://github.com/ajzbc/kanye.rest).
