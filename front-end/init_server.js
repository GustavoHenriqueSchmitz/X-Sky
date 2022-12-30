const express = require('express')
const path = require('path')
const hbs = require('hbs')

const app = express()
const port = process.env.PORT || 3000

// Define paths for Express conifg
const publicDir = path.join(__dirname, '/public')
const fontsDir = path.join(__dirname, '/fonts')
const templatesPath = path.join(__dirname, '/templates')
const partialsPath = path.join(__dirname, '/templates')

// Setup hbs engine and views location
app.set('view engine', 'hbs')
app.set('views', templatesPath)
hbs.registerPartials(partialsPath)

// Setup static directory to serve
app.use(express.static(publicDir))
app.use(express.static(fontsDir))

// Routes
app.get('/', (req, res) => {
    res.render('index')
})

app.get('/login', (req, res) => {
    res.render('login/login')
})

app.get('/sign_up', (req, res) => {
    res.render('login/registration')
})

app.listen(port, () => {
    console.log('Server is Up on port ' + port)
})
