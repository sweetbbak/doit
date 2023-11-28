default:
    go build -ldflags "-s -w"
    sudo chown root:root doit
    sudo chmod u+s doit
