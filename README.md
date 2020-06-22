FaceBook Login ( basics )

1. Setup facebook app. (developers.facebook.com )
    a.) select FaceBook Login   
    b.) Give local host that we are going to develop. e.g) https://localhost/ 
    c.) Provide OAuth releated callback API hooks (for next steps not mandatory to start with)
        e.g) return URL : https://localhost/something ( ref code)

2. configure test users. in dev mode facebook mandates developers to use test users 
    a.) go to developers.facebook.com
    b.) select App that we created
    c.) dig through left side menu for Test User configuration. 
    d.) Create & Edit User to setup passwords

3. local environment we have to create HTTPS certificate. 

    ```openssl req -x509 -out localhost.crt -keyout localhost.key \\n  -newkey rsa:2048 -nodes -sha256 \\n  -subj '/CN=localhost' -extensions EXT -config <( \\n   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")```

4. Configure APP ID in the index.js file. 

5. Run the Facebook-Auth App
```
go run main.go
```

6. Open browser https://localhost/index.html.  to see FaceBook Login button 

7. Note: first time please accept certificate


