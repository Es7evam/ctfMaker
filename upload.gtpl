<html>
    <head>
        <title>Upload file</title>

        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
        <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js" integrity="sha384-B0UglyR+jN6CkvvICOB2joaf5I4l3gm9GU6Hc1og6Ls7i6U/mkkaduKaBhlAXv9k" crossorigin="anonymous"></script>
    </head>
    <body>
        <div class="container">
            <div class="row">
                <div class="col-sm-0 col-md-0 col-lg-3 col-xl-3"></div>
                <div class="col-sm-12 col-md-12 col-lg-6 col-xl-6">
                    <form enctype="multipart/form-data" action="/uploadChall" method="post">
                        <div class="form-group">
                            <label>Challenge Name</label>
                            <input type="text" class="form-control" name="name">
                        </div>
                        <div class="form-group">
                            <label>Description</label>
                            <textarea name="desc" class="form-control" rows="5" cols="50">Sample Description</textarea>
                        </div>
                        <div class="form-group">
                            <label>Flag:</label>
                            <input type="text" class="form-control" name="flag">
                        </div>
                        <div class="form-group">
                            <label>Value</label>
                            <input type="number" class="form-control" name="value" min="0">
                        </div>
                        <div class="form-group">
                            <label>Category</label>
                            <select name="category" class="form-control">
                                {{range .Types}}
                                    <option value="{{.Value}}">{{.Name}}</option>
                                {{end}}
                            </select>
                        </div>
                        <div class="form-group">
                            <label>File</label>
                            <input type="file" class="form-control" name="uploadfile"/>
                        </div>

                        <input type="hidden" name="token" value="{{.Token}}"/>

                        <input type="submit" class="btn btn-primary" value="Upload" />
                    </form>
                </div>
                <div class="col-sm-0 col-md-0 col-lg-3 col-xl-3"></div>
            </div>
        </div>
    </body>
</html>
