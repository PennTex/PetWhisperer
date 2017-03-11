const gulp = require('gulp');
const paths = require('../paths');

gulp.task('watch', ['serve'], function(){
  gulp.watch(paths.all, ['build']);
});