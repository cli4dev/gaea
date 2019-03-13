package vue

//PackAgeJSON .
const PackAgeJSON = `
{
  "name": "{{.projectName}}",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "serve": "vue-cli-service serve --mode dev",
    "build": "vue-cli-service build --mode prod",
    "lint": "vue-cli-service lint"
  },
  "dependencies": {
    "bootstrap": "^4.1.3",
    "element-ui": "^2.4.5",
    "jquery": "^3.3.1",
    "popper.js": "^1.14.6",
    "vue": "^2.5.17",
    "vue-router": "^3.0.2",
    "vuex": "^3.0.1",
    "nav-menu":"^1.2.9",
    "login-with-up":"^1.1.1",
    "axios":"^0.18.0",
    "qs":"^6.6.0",
    "vue-cookies": "^1.5.12"
  },
  "devDependencies": {
    "@vue/cli-plugin-babel": "^3.2.0",
    "@vue/cli-plugin-eslint": "^3.2.0",
    "@vue/cli-service": "^3.2.0",
    "babel-eslint": "^10.0.1",
    "eslint": "^5.8.0",
    "eslint-plugin-vue": "^5.0.0-0",
    "vue-template-compiler": "^2.5.17"
  },
  "eslintConfig": {
    "root": true,
    "env": {
      "node": true
    },
    "extends": [
      "plugin:vue/essential",
      "eslint:recommended"
    ],
    "rules": {},
    "parserOptions": {
      "parser": "babel-eslint"
    }
  },
  "postcss": {
    "plugins": {
      "autoprefixer": {}
    }
  },
  "browserslist": [
    "> 1%",
    "last 2 versions",
    "not ie <= 8"
  ]
}
`
