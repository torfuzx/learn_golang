import "os"

func paralleWrite(data []byte) chan err {
	res := make(chan error, 2)
	f1, err := os.Create("file1")

	if err != nil {
		res <- err
	} else {
		go func() {

			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}

	f2, err := os.Create("file2")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}
