const varDecl = ['const', 'let'];

module.exports = {
	parser: '@typescript-eslint/parser',
	plugins: ['@typescript-eslint'],
	extends: ['eslint:recommended', 'plugin:@typescript-eslint/recommended', 'plugin:import/recommended', 'plugin:import/typescript'],
	root: true,
  env: {
    browser: false,
    amd: true,
    node: true
  },
	rules: {
		'arrow-body-style': [2, 'as-needed'],
		'arrow-parens': [2, 'as-needed'],
		complexity: 2,
		'constructor-super': 2,
		curly: 2,
		eqeqeq: [2, 'always', {null: 'ignore'}],
		'guard-for-in': 2,
		'import/no-unresolved': 0,
		'import/order': [
			2,
			{
				groups: ['builtin', 'external', 'unknown', ['internal', 'parent', 'sibling', 'index'], 'object', 'type'],
				pathGroups: [
					{pattern: '*.interface*', patternOptions: {matchBase: true}, group: 'external', position: 'after'}
				],
				pathGroupsExcludedImportTypes: ['builtin'],
				'newlines-between': 'always'
			}
		],
		'max-classes-per-file': [2, 1],
		'new-parens': [2, 'always'],
		'no-bitwise': 2,
		'no-caller': 2,
		'no-cond-assign': [2, 'always'],
		'no-console': 2,
		'no-else-return': 2,
		'no-eval': 2,
		'no-invalid-this': 'off',
		'no-magic-numbers': 0,
		'no-new-wrappers': 2,
		'no-redeclare': 2,
		'no-return-await': 2,
		'no-shadow': 0,
		'no-template-curly-in-string': 2,
		'no-throw-literal': 2,
		// 'no-undef-init': 2,
		'no-unsafe-finally': 2,
		'no-unused-expressions': 2,
		'object-shorthand': 2,
		'padding-line-between-statements': [2, {blankLine: 'always', prev: varDecl, next: '*'}, {blankLine: 'any', prev: varDecl, next: varDecl}],
		'prefer-arrow-callback': [2, {allowNamedFunctions: true, allowUnboundThis: true}],
		'prefer-template': 2,
		'quote-props': [2, 'as-needed'],
		quotes: 0,
		radix: 2,
		semi: [2, 'always'],
		yoda: 2,
		'@typescript-eslint/adjacent-overload-signatures': 2,
		'@typescript-eslint/array-type': [2, {default: 'array-simple'}],
		'@typescript-eslint/consistent-type-assertions': [2, {assertionStyle: 'as', objectLiteralTypeAssertions: 'allow-as-parameter'}],
		'@typescript-eslint/consistent-type-definitions': [2, 'interface'],
		'@typescript-eslint/explicit-function-return-type': [2, {allowExpressions: true}],
		'@typescript-eslint/explicit-member-accessibility': [2, {accessibility: 'explicit', overrides: {constructors: 'off'}}],
		'@typescript-eslint/member-ordering': 2,
		'@typescript-eslint/method-signature-style': [2, 'property'],
		'@typescript-eslint/naming-convention': [
			2,
			{format: ['PascalCase'], selector: ['class', 'interface']},
			{format: ['camelCase', 'UPPER_CASE'], leadingUnderscore: 'allow', selector: 'variable'}
		],
		'@typescript-eslint/no-empty-interface': 2,
		'@typescript-eslint/no-extraneous-class': [2, {allowConstructorOnly: true, allowEmpty: true, allowStaticOnly: true}],
		'@typescript-eslint/no-invalid-this': ['error'],
		'@typescript-eslint/no-magic-numbers': [1, {ignoreEnums: true}],
		'@typescript-eslint/no-namespace': 2,
		'@typescript-eslint/no-shadow': 2,
    "@typescript-eslint/no-var-requires": 0,
		'@typescript-eslint/prefer-for-of': 2,
		'@typescript-eslint/prefer-function-type': 2,
		'@typescript-eslint/prefer-namespace-keyword': 2,
		'@typescript-eslint/quotes': [2, 'single', {avoidEscape: true}],
		'@typescript-eslint/typedef': [1, {propertyDeclaration: true}],
		'@typescript-eslint/unified-signatures': 2
	},
	overrides: [
		{
			files: ['./**/*.spec.ts'],
			rules: {
				'@typescript-eslint/no-explicit-any': 0,
				'@typescript-eslint/no-magic-numbers': 0
			}
		}
	]
};
