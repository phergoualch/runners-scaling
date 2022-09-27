echo "Building..."
go build 
echo "Installing..."
sudo mv ./runners-scaling /usr/local/bin/runners

