// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 */
contract Game {
    
    //smart contract owner
    address public scOwner;

    //maps address  to how much that account has payed
    mapping(address => uint) public currentBets;
    
    modifier onlyOwner {
        require(
            msg.sender == scOwner,
            "Only owner can call this function."
        );
        _;
    }
    
    //allow owner to deposit money to execute SC functions from server
    //This is 'payable' so that the server ('scOwner') has funds (JORDAN WANTS GAS MONEY) to execute 'chooseWinner'
    constructor() payable {
        scOwner = msg.sender; //You can hard-code this to be the server's address ('scOwner = msg.sender' is used for testing / development)
    }
    
    //player wages their bet here
    //when the client executes this function, it pays a small amount of GAS
    function applyWager() public payable {
        require(
            msg.value >= 1 ether && currentBets[msg.sender] == 0,
            "Did not wage enough ether OR already have an existing wager"
        );
        currentBets[msg.sender] = msg.value;
    }
    
    //when the server executes this function, it pays a small amount of GAS
    function chooseWinner(address winner, address loser) public onlyOwner {
        //this should prevent a re-entrancy attack. Read:
        //https://docs.soliditylang.org/en/v0.8.4/security-considerations.html#re-entrancy
        uint sum = currentBets[winner] + currentBets[loser];
        currentBets[winner] = 0;
        currentBets[loser] = 0;
        payable(winner).transfer(sum);
    }
}