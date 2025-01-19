<script>
    import { Link } from "svelte-routing";
    import Nav from "./Nav.svelte";
</script>

<main>
    <Nav />
    <h3>-> 243 ONLINE</h3>
    <h1>Share my Story</h1>

    <div class="chat">
        <div class="message message--inc">
            <p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi
                ducimus minima, deleniti sapiente suscipit corrupti libero enim
                fugit dolorem tempore mollitia sit vero alias architecto, ab
                explicabo eos a molestiae.
            </p>
            <p><b>Random User #47</b> 10PM</p>
        </div>
        <div class="message message--out">
            <p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi
                ducimus minima, deleniti sapiente suscipit corrupti libero enim
                fugit dolorem tempore mollitia sit vero alias architecto, ab
                explicabo eos a molestiae.
            </p>
            <p><b>You</b> 10PM</p>
        </div>
    </div>
    <div class="reply__box">
        <div class="reply__input">
            <input type="text" placeholder="Type something" />
        </div>
        <div class="reply__btn">
            <p>send</p>
        </div>
    </div>
</main>

<style>
    .chat {
        position: relative;
    }
    .message {
        border-radius: 1rem;
        padding: 1rem;
        width: 70%;
        margin-bottom: 2rem;
    }
    .message p:first-child {
        margin-bottom: 1rem;
    }
    .message--inc {
        background-color: var(--color-light);
    }
    .message--out {
        background-color: var(--color-primary-light);
        margin-left: 60%;
        translate: -45%;
    }
    .reply__box {
        display: flex;
    }
    .reply__input {
        flex: 1;
    }
    .reply__input input {
        border-radius: 4rem;
    }
    .reply__btn {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0.5rem;
    }
</style>
