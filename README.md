# golang-face-detector

A simple and fast face detector using gocv and har cascade classifier.

## Usage
Need to install opencv and gocv in your machine
---
- install gocv and open cv  using the following commands.

    ```go get -u -d gocv.io/x/gocv```

    ```cd $GOPATH/src/gocv.io/x/gocv```

    ```make install```
- If it works correctly, at the end of the entire process, the following message should be displayed:

    ```gocv version: 0.17.0```

    ``` opencv lib version: 4.0.0```
- Install gtk and gtk3 modules. 

    ```sudo apt install libcanberra-gtk-module libcanberra-gtk3-module```
    ___
- To run the program 

    ```go run main.go -image input/image.jpg```