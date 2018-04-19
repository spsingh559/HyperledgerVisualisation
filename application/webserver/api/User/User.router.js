const router = require('express').Router();


const userController = require('./User.controller.js');
router.post('/',userController.loginData);
router.post('/signUp',userController.signUpData);


exports = module.exports = router;
