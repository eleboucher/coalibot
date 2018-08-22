module.exports = {
  apps: [
    {
      name: 'Coalibot',
      script: './src/index.js',
      env: {
        NODE_ENV: 'development'
      },
      watch: true,
      ignore_watch: '*.json',
      env_production: {
        NODE_ENV: 'production'
      }
    }
  ]
}
