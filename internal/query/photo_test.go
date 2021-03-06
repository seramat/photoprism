package query

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/form"
)

/*func TestQuery_Photos(t *testing.T) {
	conf := config.TestConfig()

	search := New(conf.OriginalsPath(), conf.Db())

	t.Run("search with query", func(t *testing.T) {
		query := form.NewPhotoSearch("Title:Reunion")
		result, err := search.Photos(query)

		t.Log(result)
		t.Log(err)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, "Cake", result[1].PhotoName)
		assert.Equal(t, "COW", result[0].PhotoName)
	})
}*/

func TestQuery_PhotoByID(t *testing.T) {
	conf := config.TestConfig()

	search := New(conf.Db())

	t.Run("photo found", func(t *testing.T) {
		result, err := search.PhotoByID(1000000)
		assert.Nil(t, err)
		assert.Equal(t, 2790, result.PhotoYear)
	})

	t.Run("no photo found", func(t *testing.T) {
		result, err := search.PhotoByID(99999)
		assert.Error(t, err, "record not found")
		t.Log(result)
	})
}

func TestQuery_PhotoByUUID(t *testing.T) {
	conf := config.TestConfig()

	search := New(conf.Db())

	t.Run("photo found", func(t *testing.T) {
		result, err := search.PhotoByUUID("pt9jtdre2lvl0y12")
		assert.Nil(t, err)
		assert.Equal(t, "Reunion", result.PhotoTitle)
	})

	t.Run("no photo found", func(t *testing.T) {
		result, err := search.PhotoByUUID("99999")
		assert.Error(t, err, "record not found")
		t.Log(result)
	})
}

func TestQuery_PreloadPhotoByUUID(t *testing.T) {
	conf := config.TestConfig()

	search := New(conf.Db())

	t.Run("photo found", func(t *testing.T) {
		result, err := search.PreloadPhotoByUUID("pt9jtdre2lvl0y12")
		assert.Nil(t, err)
		assert.Equal(t, "Reunion", result.PhotoTitle)
	})

	t.Run("no photo found", func(t *testing.T) {
		result, err := search.PreloadPhotoByUUID("99999")
		assert.Error(t, err, "record not found")
		t.Log(result)
	})
}

func TestSearch_Photos(t *testing.T) {
	conf := config.TestConfig()

	conf.CreateDirectories()

	search := New(conf.Db())

	t.Run("normal query", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("label query", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "label:dog"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			// TODO: Add database fixtures to avoid failing queries
			t.Logf("query failed: %s", err.Error())
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("invalid label query", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "label:xxx"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		assert.Equal(t, err.Error(), "label \"xxx\" not found")

		t.Logf("results: %+v", photos)
	})
	t.Run("form.location true", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 3
		f.Offset = 0
		f.Location = true

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.camera", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 3
		f.Offset = 0
		f.Camera = 2

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.color", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 3
		f.Offset = 0
		f.Color = "blue"

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.favorites", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "favorites:true"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.country", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "country:de"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.title", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "title:Pug Dog"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.hash", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "hash:xxx"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.duplicate", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "duplicate:true"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.portrait", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "portrait:true"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.mono", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "mono:true"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.chroma", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "chroma:50"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.fmin and Order:oldest", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Fmin:5 Order:oldest"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.fmax and Order:newest", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Fmax:2 Order:newest"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.Lat and form.Lng and Order:imported", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Lat:33.45343166666667 Lng:25.764711666666667 Dist:2000 Order:imported"
		f.Count = 3
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})
	t.Run("form.Before and form.After", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Before:2005-01-01 After:2003-01-01"
		f.Count = 5000
		f.Offset = 0

		photos, _, err := search.Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", photos)
	})

}
