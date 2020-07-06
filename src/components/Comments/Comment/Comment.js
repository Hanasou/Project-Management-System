import React from 'react';

import Wrapper from '../../../hoc/Wrapper';

const comment = (props) => {
    const email = props.email;
    const content = props.children;
    return(
        <Wrapper>
            <li>{email}<br/>{content}</li>
        </Wrapper>
    )
}

export default comment;