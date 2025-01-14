/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

module.exports = {
  globals: {
    window: true,
  },
  setupFiles: [
    '<rootDir>/src/testing/enzyme-setup.ts',
    'jest-canvas-mock',
  ],
  setupFilesAfterEnv: [
    '<rootDir>/src/testing/jest-test-setup.js',
  ],
  moduleFileExtensions: [
    'js',
    'json',
    'jsx',
    'mjs',
    'ts',
    'tsx',
  ],
  moduleDirectories: [
    'node_modules',
    '<rootDir>/src',
    '<rootDir>/packages/pixie-components/src',
    '<rootDir>/packages/pixie-api/src',
    '<rootDir>/packages/pixie-api-react/src',
  ],
  moduleNameMapper: {
    '^.+.(jpg|jpeg|png|gif|svg)$': '<rootDir>/src/testing/file-mock.js',
    '(\\.css|\\.scss$)|(normalize.css/normalize)|(^typeface)|(^exports-loader)': 'identity-obj-proxy',
    'monaco-editor': '<rootDir>/node_modules/react-monaco-editor',
  },
  resolver: null,
  transform: {
    '^.+\\.jsx?$': 'babel-jest',
    '^.+\\.tsx?$': 'ts-jest',
    '^.+\\.toml$': 'jest-raw-loader',
  },
  testRegex: '.*test\\.(ts|tsx|js|jsx)$',
  reporters: [
    'default',
    'jest-junit',
  ],
  collectCoverageFrom: [
    'src/**/*.ts',
    'src/**/*.tsx',
    'src/**/*.js',
    'src/**/*.jsx',
    'src/*.ts',
    'src/*.tsx',
    'src/*.js',
    'src/*.jsx',
  ],
};
