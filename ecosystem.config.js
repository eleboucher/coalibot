module.exports = {
  apps: [
    {
      name: 'Coalibot',
      script: './src/index.js',
      env: {
        NODE_ENV: 'development'
      },
      watch: './src/',
      env_production: {
        NODE_ENV: 'production'
      }
    }
  ]
}
