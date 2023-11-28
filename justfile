a1 := "env"

default:
    go build -ldflags "-s -w"
    sudo chown root:root doit
    sudo chmod u+s doit
allows:
    go build -ldflags "-s -w -X main.build_allow=true -X main.allow5=env"
    sudo chown root:root doit
    sudo chmod u+s doit
onetime:
    go build -ldflags "-s -w -X main.delete=true"
    sudo chown root:root doit
    sudo chmod u+s doit
