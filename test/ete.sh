execc=exec

run(){
  docker run -dit --name ete cbot918/ubugo
  docker $execc ete bash -c "git clone https://github.com/cbot918/ya && cd ya && go build ."
  docker $execc ete bash -c "mv ya/ya /bin"
  docker $execc -it ete bash
}
run 
once(){
  echo hi
}