<html>
<head>
       <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="/uploadChall" method="post">
    Challenge Name:<input type="text" name="name"> <br>
    Description:
    <textarea name="desc" rows="5" cols="50"> Sample Description </textarea> <br>
    Flag:<input type="text" name="flag"> <br>
    Value:<input type="number" name="value" min="0"> <br>

    Category: 
    <select name="category">
    <option value="cat1">apple</option>
    <option value="cat2">pear</option>
    <option value="cat3">banana</option>
    </select> <br>

    File: 
    <input type="file" name="uploadfile" />
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload" />
</form>
</body>
</html>