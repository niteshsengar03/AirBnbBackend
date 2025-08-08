# Steps to setup the template

1. Clone the project 
using http
```
git clone https://github.com/niteshsengar03/express-typescript-template.git <Project name>
```
or ssh
```
git clone git@github.com:niteshsengar03/express-typescript-template.git <Project name>
```

2. move into the folder structure
```
cd <Project name>
```

3. Install npm dependencies
```
npm i
```

4. Create a new .env file in the root directory and add the `Port` env variable
```
echo PORT=3000 >> .env
```

5. Start the express server
```
npm run dev
```