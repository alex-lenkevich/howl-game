package game

import (
	"testing"

	"reflect"
)

func TestConnect(t *testing.T) {

	t.Run("Save/Load Stats", func(t *testing.T) {
		connection, err := Connect("mongodb://howl:howl@localhost:27017/howl")
		if err != nil {
			t.Errorf("Connect() error = %v\n", err)
			return
		}
		statsEtalon := Stats{12, 42, 3, 8}
		err = connection.SaveStats(statsEtalon)
		if err != nil {
			t.Errorf("SaveStats() error = %v", err)
			return
		}
		stats, err := connection.LoadStats(12)
		if err != nil {
			t.Errorf("LoadStats() error = %v", err)
			return
		}

		if !reflect.DeepEqual(stats, statsEtalon) {
			t.Errorf("want %+v, got %+v", statsEtalon, stats)
		}
	})


}
