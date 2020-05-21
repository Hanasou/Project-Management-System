import React from 'react';
import './App.css';

import Wrapper from './hoc/Wrapper'
import Auth from './containers/Auth/Auth'
function App() {
  
  return (
    <Wrapper>
      <Auth></Auth>
    </Wrapper>
  );
}

export default App;
