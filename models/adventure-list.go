package models

import "time"

// Jen is the original creator of all adventures
var Jen = User{
	ID:       1,
	Name:     "Jennifer Petersen",
	Gender:   "Female",
	Age:      31,
	JoinedAt: time.Date(2012, time.June, 23, 9, 33, 0, 0, time.UTC),
}

// AdventureTests are a list of premade adventures for testing
var AdventureTests = Adventures{
	{
		ID:          1,
		Title:       "The Oregon Zoo",
		Body:        "This is the body",
		Description: "Why not Portland? Beautiful scenery, amazing urban markets. Not to mention the drive down. Count me in!",
		PlannedDate: time.Date(2016, time.March, 13, 20, 0, 0, 0, time.UTC),
		Completed:   false,
		Author:      &Jen,
	},
	{
		ID:          2,
		Title:       "Niagara Falls",
		Body:        "This is the body",
		Description: "Beautiful scenery, roaring waters. I can't even begin to imagine how amazing this place is in person. Not to mention that at night the waters are illuminated with different colors.",
		PlannedDate: time.Date(2017, time.July, 15, 16, 0, 0, 0, time.UTC),
		Completed:   false,
		Author:      &Jen,
	},
	{
		ID:          3,
		Title:       "COSTCO",
		Body:        "This is the body",
		Description: "Now I know what you're thinking; \"Coscto!? How is Costco an adventure?\" Well, it's simple. Free samples, rows and rows of mile high aisles. Oh and did I mention the samples? Every trip is different. You never know what you may find.",
		PlannedDate: time.Now(),
		Completed:   false,
		Author:      &Jen,
	},
	{
		ID:          4,
		Title:       "San Francisco",
		Body:        "This is the body",
		Description: "I don't know nothing about it nor have I been there. When I think of this city, I think of Rice a Roni and Full House. They both are great so I'm sure the city is too. Oh and an Alcatraz tour is a must.",
		PlannedDate: time.Date(2016, time.August, 18, 14, 0, 0, 0, time.UTC),
		Completed:   false,
		Author:      &Jen,
	},
	{
		ID:          5,
		Title:       "FORKS, WASHINGTON",
		Body:        "This is the body",
		Description: "Who wouldn't want to visit the hometown of Bell and Edward Cullen? Who doesn't love sparkly vampires? We could take pictures by the sign and by Bella's actual truck. Great, right?",
		PlannedDate: time.Date(2016, time.February, 20, 12, 0, 0, 0, time.UTC),
		Completed:   false,
		Author:      &Jen,
	},
	{
		ID:          6,
		Title:       "Bora Bora",
		Body:        "This is the body",
		Description: "I think this picture says it all. I don't need to say anything else. Just you and me, matching bandaids from our Malaria vaccines and the big net over our bed.",
		PlannedDate: time.Date(2018, time.July, 15, 16, 0, 0, 0, time.UTC),
		Completed:   false,
		Author:      &Jen,
	},
	{
		ID:          7,
		Title:       "Tillamook Cheese Factory",
		Body:        "This is the body",
		Description: "Who doesn't love cheese? You know how I feel about samples. Oh and I almost forgot that it's down Hwy 101, a scenic coastal hwy down the coast of Washington and Oregon.",
		PlannedDate: time.Date(2016, time.April, 12, 20, 0, 0, 0, time.UTC),
		Completed:   false,
		Author:      &Jen,
	},
}
