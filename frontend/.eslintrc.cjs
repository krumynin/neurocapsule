module.exports = {
  root: true,
  extends: '@nuxtjs/eslint-config-typescript',
  rules: {
    semi: ['error', 'always'],
    curly: ['error', 'multi-line'],
    'import/order': 'off',
    'operator-linebreak': 'off',
    camelcase: ['error', {
      properties: 'never',
      ignoreDestructuring: false,
    }],
    'comma-dangle': ['error', {
      arrays: 'always-multiline',
      objects: 'always-multiline',
      imports: 'always-multiline',
      exports: 'always-multiline',
      functions: 'always-multiline',
    }],
    '@typescript-eslint/no-unused-vars': ['error', {
      vars: 'all',
      args: 'after-used',
      ignoreRestSiblings: true,
    }],
    'require-await': 'off',
    'vue/multi-word-component-names': 'off',
  },
};
