import { globalIgnores } from 'eslint/config'
import { defineConfigWithVueTs, vueTsConfigs } from '@vue/eslint-config-typescript'
import pluginVue from 'eslint-plugin-vue'
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'

const autoImportConfig = require('./.eslintrc-auto-import.json')

// To allow more languages other than `ts` in `.vue` files, uncomment the following lines:
// import { configureVueProject } from '@vue/eslint-config-typescript'
// configureVueProject({ scriptLangs: ['ts', 'tsx'] })
// More info at https://github.com/vuejs/eslint-config-typescript/#advanced-setup

export default defineConfigWithVueTs(
  {
    name: 'app/files-to-lint',
    files: ['**/*.{ts,mts,tsx,vue}'],
  },

  globalIgnores(['**/dist/**', '**/dist-ssr/**', '**/coverage/**']),

  pluginVue.configs['flat/essential'],
  vueTsConfigs.recommended,
  skipFormatting,
  {
    languageOptions: {
      globals: {
        autoImportConfig
      }
    }
  },
  {
    rules: {
      'import/no-unresolved': [0],
      'vue/multi-word-component-names': 'off',
      'vue/no-deprecated-router-link-tag-prop': 'off',
      'import/extensions': 'off',
      'import/no-absolute-path': 'off',
      'no-async-promise-executor': 'off',
      'import/no-extraneous-dependencies': 'off',
      'vue/no-multiple-template-root': 'off',
      'vue/html-self-closing': 'off',
      'no-console': 'off',
      'no-plusplus': 'off',
      'no-useless-escape': 'off',
      'no-bitwise': 'off',
      '@typescript-eslint/no-explicit-any': ['off'],
      '@typescript-eslint/explicit-module-boundary-types': ['off'],
      '@typescript-eslint/ban-ts-comment': ['off'],
      'vue/no-setup-props-destructure': ['off'],
      '@typescript-eslint/no-empty-function': ['off'],
      'vue/script-setup-uses-vars': ['off'],
      //can config  to 2 if need more then required
      '@typescript-eslint/no-unused-vars': [0],
      'no-param-reassign': ['off'],
      '@typescript-eslint/no-extra-non-null-assertion': 'off',
      '@typescript-eslint/no-non-null-assertion': 'off',
    }
  }
)
