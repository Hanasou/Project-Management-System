import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import Wrapper from './hoc/Wrapper'
import Signup from './containers/Auth/Signup'
function App() {
  
  return (
    <Wrapper>
      <Signup/>
    </Wrapper>
  );
}

export default App;
