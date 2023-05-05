package database

const (
	AddPlatformQuery = `INSERT INTO video_games.platform (platform_name)
						  VALUES ((?));`

	DeletePlatformQuery1 = `
DELETE FROM video_games.region_sales WHERE game_platform_id IN 
                                           (SELECT id FROM video_games.game_platform WHERE platform_id IN
                                                                                           (SELECT id FROM video_games.platform WHERE platform_name =?));
		`
	DeletePlatformQuery2 = `DELETE FROM video_games.game_platform WHERE platform_id IN
						                                              (SELECT id FROM video_games.platform WHERE platform_name =?);
			`
	DeletePlatformQuery3 = `DELETE FROM video_games.platform WHERE platform_name = ?;`
)
