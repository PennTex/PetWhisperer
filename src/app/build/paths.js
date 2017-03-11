const appRoot = 'src/';
const outputRoot = 'dist/';

module.exports = {
  all: [appRoot + "**/*.js", appRoot + "**/*.html"],
	root: appRoot,
	html: [
		appRoot + "**/*.html"
	],
  app: appRoot + "main.js",
  output: outputRoot
}