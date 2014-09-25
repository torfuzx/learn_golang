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
  
  ```
  export PATH=$PATH:/usr/local/go/bin
  ```

