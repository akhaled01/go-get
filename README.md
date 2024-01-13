# gopher-get

This is a lightweight basic version of GNU `wget` written purely in <a href="https://go.dev/doc/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/skills/go-colored.svg" width="200" height="200" alt="Go" /></a>

## Implemented Options

1. `-B` : Enables Silent Mode. All output will be written to a file called `wget-log.txt`.
2. `-O` and `-P` : rename the file under a different name and under a different path respectively.
3. The project implements a rate limiter (still in works). Basically the program can control the speed of the download by using the flag `--rate-limit`. If you download a huge file you can limit the speed of your download, preventing the program from using the full possible bandwidth of your connection.
4. Downloading different files is possible. For this the program will receive the `-i` flag followed by a file name that will contain all links that are to be downloaded. The downloads will be done in async.
5. Finally, the project is able to mirror a website using the `-mirror` tag (in works).

## More Information

* The project is written in pure golang, with a makefile for creating a build by running the `make` command.
* Multiple different external repositories were used, such as `github.com/progressbar/v3` for progress bar functionality.
* The project tried to use as much of the stdlibs as possible, but had to resort to external packages for some functionality like HTML parsing.

### Written By akhaled
