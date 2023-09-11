ALTER TABLE post_approval DISABLE TRIGGER user_rating_trigger_1;

ALTER TABLE comment_approval DISABLE TRIGGER user_rating_trigger_2;

ALTER TABLE post_edition DISABLE TRIGGER user_rating_trigger_3;

ALTER TABLE post_approval DISABLE TRIGGER post_rating_calculating_trigger;