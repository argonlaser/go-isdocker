[![GoDoc](https://godoc.org/github.com/argonlaser/go-isdocker?status.svg)](http://godoc.org/github.com/argonlaser/go-isdocker)

# go-isdocker

  Check if the process is running inside a docker.

## Getting Started

This package checks for the presence of ```/.dockerenv``` file or searches for the presence of ```docker``` in the file ```/proc/self/cgroup``` 

```
import "github.com/argonlaser/go-isdocker"
```

## Documentation and APIs
Refer godoc [go-isdocker](https://godoc.org/github.com/argonlaser/go-isdocker) for detailed APIs.

### Usage

```
import "github.com/argonlaser/go-isdocker"

if IsDocker() == true {
  // this process is running on a docker

} else {
  // this process is running on a host machine

}
```

## Running the tests

```
$ cd $HOME/go/src/github.com/argonlaser/go-isdocker
$ go test
```

### Check coding style tests

```
$ cd $HOME/go/src/github.com/argonlaser/go-isdocker
$ go fmt
```

## Bug reports
Please feel to raise an [issue](https://github.com/argonlaser/go-isdocker/issues) here. It means a lot to the project.


## Contributing
1. Fetch the Sources From GitHub

```
$ go get github.com/argonlaser/go-isdocker
```

2. Change to the Hugo source directory:

```
$ cd $GOPATH/src/github.com/argonlaser.go-isdocker
```

3. Create a new branch for your changes (the branch name is arbitrary):

```
$ git checkout -b mybranch
```

4. After making your changes, commit them to your new branch:
 
```
$ git commit -a -v
```

5. Fork this project in Github and add your fork as a new remote (the remote name, "fork" in this example, is arbitrary):

```
$ git remote add fork git://github.com/USERNAME/go-isdcoker.git
```

6. Run codestyle checks and tests for ```go-isdocker``` with Your Changes. Add more tests if needed.

```
$ cd $HOME/go/src/github.com/argonlaser/go-isdocker
$ go fmt
$ go test
```

7. Push the changes to your new remote once all tests are passed.

```
$ git push --set-upstream fork mybranch
```

## Authors

* Authored and maintained by **Venkata krishna** and the list of [contributors](https://github.com/argonlaser/go-isdocker/contributors).

Contact me in | [Twitter](https://twitter.com/argon_laser) | [Gmail](vkvenkat94@gmail.com) |

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/argonlaser/go-isdocker/blob/master/LICENSE) file for details

## Acknowledgments

* [Hugo-Readme](https://github.com/spf13/hugo/blob/master/CONTRIBUTING.md )
