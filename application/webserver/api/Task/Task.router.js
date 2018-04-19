const router = require('express').Router();
 const TaskController = require('./Task.controller.js');
router.get('/:timeStamp', TaskController.getTaskDetail);
router.get('/searchName/:searchName',TaskController.getSearchResult);
router.post('/',TaskController.postData);
router.put('/',TaskController.updateTask);
router.delete('/:_id',TaskController.removeTask);
exports = module.exports = router;
