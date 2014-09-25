Install golang 1.2.2 on Ubuntu 14.04
------------------------------------


1. Find the go release tarball at the offical downbload page: http://golang.org/dl/


2. Download the tarball 

  using wget:
  ```bash 
  cd /tmp
  wget https://storage.googleapis.com/golang/go1.2.2.linux-amd64.tar.gz
  ```

  or using curl: 
  ```bash
  curl -O https://storage.googleapis.com/golang/go1.2.2.linux-amd64.tar.gz
  ```

3. Extract the tarball to `/usr/local`

  ```bash
  sudo tar -C /usr/local -xvzf go1.2.2.linux-amd64.tar.gz
  ```

4. Make the go's binaries dicoverable by the system 
  
  Add this line to the `~/.bashrc`:
  ```
  export PATH=$PATH:/usr/local/go/bin
  ```
  and then apply the new change:
  
  ```bash
  source ~/.bashrc
  ```
5. Make sure go's installed, by running `go version`
  
  You should see the following output: 
  ```
  $ go version
  go version go1.2.2 linux/amd64
  ```

6. Set up the GOPATH environment variable 

  Choose a folder as you gopath, here I choose '/home/kurt/gopath' as mine, and add the following lines to the `./bashrc`
  ```
  export GOPATH="/home/kurt/gopath"
  export PATH=$PATH:/home/kurt/gopath/bin
  ```

  
  
