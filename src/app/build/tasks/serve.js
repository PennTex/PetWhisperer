const gulp      = require('gulp'),
  browserSync = require('browser-sync'),
  paths       = require('../paths.js');

gulp.task('serve', ['build'], function(done) {
  browserSync({
    online: false,
    open: false,
    port: 5001,
    server: {
      baseDir: [paths.output],
      middleware: function(req, res, next) {
        res.setHeader('Access-Control-Allow-Origin', '*');
        next();
      }
    }
  }, done);
});