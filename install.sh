install(){
  go build .
  sudo mkdir -p /usr/local/bin/ya && sudo mv ya /usr/local/bin/ya
  echo "
export PATH="\$PATH:/usr/local/bin/ya"  
" >> ~/.bashrc
  echo "\n ya installed \n"
}
install