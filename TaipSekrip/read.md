# Setup Typescript Environment

## Setup Project
* npm init -y or npm init
* add type module to package json

## Setup Unit Test
* npm install --save-dev jest @types/jest
* npm install --save-dev babel-jest @babel/preset-env
* create babel.config.json
* check here: https://babeljs.io/setup#installation

## Setup TypeScript Project
* npm install --save-dev typescript
* npx tsc --init
* set tsconfig.json module from common js to es6 if want use es6
* https://jestjs.io/docs/getting-started#using-typescript
* npm install --save-dev @babel/preset-typescript and update babel.config.json
* npm install --save-dev ts-jest 
* npm install --save-dev @jest/globals
* npm install --save-dev @types/jest