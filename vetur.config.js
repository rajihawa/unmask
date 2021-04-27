module.exports = {
    projects: [
      './web', // shorthand for only root.
      {
        // **required**
        // Where is your project?
        // It is relative to `vetur.config.js`.
        root: './web',
        // **optional** default: `'package.json'`
        // Where is `package.json` in the project?
        // We use it to determine the version of vue.
        // It is relative to root property.
        package: './web/package.json',
        // **optional**
        // Where is TypeScript config file in the project?
        // It is relative to root property.
        tsconfig: './web/tsconfig.json',
      }
    ]
  }