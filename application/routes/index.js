module.exports = function(app) {

   app.all("/login",function(req, res, next) {
        res.render('login');

    });
   /* app.all("/userLogin",function(req, res, next) {
        res.render('login');

    });
  */
    
    app.all("/customer_index",function(req, res, next) {
        res.render('customer_index');
    });

        app.all("/transporter_index",function(req, res, next) {
        res.render('transporter_index');
    });
        app.all("/uniper_index",function(req, res, next) {
        res.render('uniper_index');
    });
       app.all("/supplier_index",function(req, res, next) {
        res.render('supplier_index');
    });

    app.all("/",function(req, res, next) {
        res.render('login');
    });
}
