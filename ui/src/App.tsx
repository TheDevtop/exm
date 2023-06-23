import * as React from 'react'
import * as Material from '@mui/material/'
import MenuBar from './MenuBar'
import * as Lib from './lib'

export default function App() {
  const [rex, setRex] = React.useState('');
  const [obj, setObj] = React.useState('');

  function clicked() {
    /*
    Use the lib package to communicate with EXM.
    Print the results as a table in the result card.
    */
   const rf = Lib.Search(rex, obj);
   console.log(rf);
  }

  return (
    <Material.Container fixed>
      <MenuBar />
      <Material.Grid container spacing={4}>
        <Material.Grid item xs={6}>
          <Material.Card><Material.CardContent>
          <Material.TextField margin="dense" onChange={r => setRex(r.target.value)} id="outlined-basic" label="Regex" variant="outlined" />
          <Material.TextField margin="dense" onChange={o => setObj(o.target.value)} id="outlined-basic" label="Object" variant="outlined" />
          </Material.CardContent>
          <Material.CardActions>
            <Material.Button onClick={clicked} variant="contained">Search</Material.Button>
          </Material.CardActions>
          </Material.Card>
        </Material.Grid>
        <Material.Grid item xs={6}>
          <Material.Card><Material.CardContent>
              Results
          </Material.CardContent></Material.Card>
        </Material.Grid>
      </Material.Grid>
    </Material.Container>
  );
}
