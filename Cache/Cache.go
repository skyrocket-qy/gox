package Cache

import (
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/skyrocketOoO/GoUtils/Reflex"
)

type Cache[K cmap.Stringer, V any] struct {
	cmap.ConcurrentMap[K, V]
	database     *gorm.DB
	primaryField string
}

func NewCache[K cmap.Stringer, V any](database *gorm.DB, autoLoad bool) (out *Cache[K, V]) {
	primaryField := ""
	fields := dao.ParsePrimaryFields[V](database)

	if 1 == len(fields) {
		primaryField = fields[0]
	}

	out = &Cache[K, V]{
		ConcurrentMap: cmap.NewStringer[K, V](),
		database:      database,
		primaryField:  primaryField,
	}

	if autoLoad {
		_ = out.Load()
	}

	return
}

func GetCache[K cmap.Stringer, V any](
	data *Cache[K, V],
	database *gorm.DB,
	autoLoad bool,
) (out *Cache[K, V]) {
	if nil != data {
		return
	}

	out = NewCache[K, V](database, autoLoad)
	return
}

func (it *Cache[K, V]) ReadItem(key K) (out V, throw error) {
	item, throw := dao.GetItem[V](it.database.Session(dao.SilentSession), key)
	if nil != throw {
		return
	}

	it.Set(key, *item)
	return
}

func (it *Cache[K, V]) GetItem(key K) (out V, throw error) {
	out, ok := it.ConcurrentMap.Get(key)
	if ok {
		return
	}

	out, throw = it.ReadItem(key)
	return
}

func (it *Cache[K, V]) HasItem(key K) (ok bool, throw error) {
	ok = it.ConcurrentMap.Has(key)
	if ok {
		return
	}

	// ok, throw = dao.HasItem[V](it.database.Session(dao.SilentSession), key)
	ok, throw = dao.HasItem[V](it.database, key)
	if nil != throw {
		ok = false
	}

	return
}

func (it *Cache[K, V]) GetValues() (out []V) {
	out = lo.Values(it.Items())
	return
}

func (it *Cache[K, V]) Load() (throw error) {
	var items []V
	it.database.FindInBatches(
		&items,
		dao.ReadBatchSize,
		func(database *gorm.DB, batch int) (throw error) {
			for _, item := range items {
				var key *K
				key, throw = Reflex.GetField[K](&item, it.primaryField)
				if nil != throw {
					return
				}

				it.Set(*key, item)
			}

			return
		},
	)

	return
}
