package mongo

type productsRepository struct {
	db      string
	session *mongo.Session
}

func NewProductsRepository(db string, session *mongo.Session) (db.ProductsRepository, error) {
	r := &productsRepository{
		db:      db,
		session: session,
	}

	index := mongo.Index{
		Key:        []string{"trackingid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("cargo")

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}

	return r, nil

	return mongo.WithSession(context.TODO(), session, func(sessCtx mongo.SessionContext) error {
		// Use sessCtx as the Context parameter for InsertOne and FindOne so both operations are run under the new
		// Session.

		coll := client.Database(db).Collection("coll")
		res, err := coll.InsertOne(sessCtx, bson.D{{"x", 1}})
		if err != nil {
			return err
		}
		fmt.Println("gg")

		var result bson.M
		gg := coll.FindOne(sessCtx, bson.D{{"_id", res.InsertedID}})
		if err = gg.Decode(&result); err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	})
}

func New(url string, timeout string) (*mongo.Client, *mongo.Session, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}

	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := client.StartSession(opts)
	if err != nil {
		return nil, nil, err
	}
	// defer sess.EndSession(context.TODO())

	return client, sess, nil
}
