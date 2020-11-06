-- +migrate Up
CREATE TABLE IF NOT EXISTS `email_activations` (
  `activation_token` VARCHAR(255) NOT NULL PRIMARY KEY,
  `user_id` VARCHAR(255) NOT NULL,
  `expires_at` INT NOT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS `email_activations`;