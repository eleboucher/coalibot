module.exports = {
  apps: [
    {
      name: "Coalibot",
      script: "./src/index.js",
      env: {
        NODE_ENV: "production"
      },
      watch: "./src/"
    }
  ]
};
