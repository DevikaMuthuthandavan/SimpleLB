# SimpleLB
A Simple Load Balancer using go lang. 

*To create a server in Python:*
python -m venv pyvenv --Create a virtual env
source pyvenv/bin/activate -- To activate the created venv
pip install flask -- Install flask
python3 server.py  -- To run the server

-------------------------------------------------------------------------------------------------------------------
*To run Load Balancer:*
go run main.go

-------------------------------------------------------------------------------------------------------------------
*To run server.py*
source pyvenv/bin/activate
python3 server.py "ServerName" "portNo" --Eg., "Server 1" "5001"

-------------------------------------------------------------------------------------------------------------------
*Load test with curl:*
for i in {1..10}; do curl 127.0.0.1:8000; done