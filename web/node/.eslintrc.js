module.exports = {
	root: true,
	env: {
		es6: true,
		node: true,
		amd: true,
		jquery: true,
		commonjs: true,
		jasmine: true
	},
	globals: {
		$: false,
		_: false,
		define: true,
		module: false,
		jQuery: false,
		lodash: false,
		inject: false,
		require: false
	},
	plugins: [
		'lodash'
	],
	extends: ['plugin:lodash/recommended']
};
