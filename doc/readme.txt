bee api peony -tables="equipment_gateway,equipment_simcards" -driver=mysql -conn="root:root@tcp(127.0.0.1:3306)/gdkxdl"

生成swagger文档
bee run -gendoc=true -downdoc=true


生成RSA加密算法
openssl genrsa -out app.rsa keysize
openssl rsa -in app.rsa -pubout > app.rsa.pub
可以在以下网站进行验证
https://jwt.io/


