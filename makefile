create_ssh_key:
	ssh-keygen -t rsa
	ssh-keygen -f id_rsa.pub -e -m pem > id_rsa.pub.pem
	ssh-keygen -p -m PEM -f id_rsa

	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub

	ssh-keygen -b 4096 -t rsa
