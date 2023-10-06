---
title: "GnoMobile, a Framework for Building Gno Mobile Apps"
publication_date: 2023-09-29T13:37:00Z
slug: gnomobile
tags: [gnomobile, berty, wesh network]
authors: [jeff, costin, remi, iuri]
---

*This blog post is written by Berty Technologies, an NGO that is building open and free communication solutions without any of the limitations imposed by centralized systems. Berty is a proud partner and grantee of Gno.land.*

The year is 2023. Current Gno apps run on desktop or laptop computers that have Go installed. To run on mobile, the app would need to bundle the Go runtime, which is complicated for most developers. At Berty, we have years of experience using Go on mobile and overcoming difficulties with Android and iOS operating systems. We built Wesh Network, a decentralized communication protocol that enables p2p users to reliably and securely send messages over async networks, even in environments with poor or no connectivity.

This stage is thus set to take the leap and make it easier for builders to develop Gno applications for mobile devices.

# What is GnoMobile?

Simply put, GnoMobile is a framework for developing Gno mobile applications. This is how it works:

*WARNING: Deep technical sections ahead. Grab a coffee before venturing forth*.

For communication between the mobile app and the Gno code, GnoMobile uses [gRPC](https://grpc.io/), a well-supported framework that sends and receives Google Protobuf messages. Even though the core Gno code is written in Go, the app code can use React Native, Java, Swift, etc. The following system diagram shows how gRPC is used.

<div align="center">
    <img style="display: block;-webkit-user-select: none;margin: auto;cursor: zoom-in;background-color: hsl(0, 0%, 90%);transition: background-color 300ms;" src="https://github-production-user-asset-6210df.s3.amazonaws.com/109347079/267934754-e4da6fec-a586-4ebe-97cc-3b3ad7f79370.jpg" width="324" height="300" align="left">
</div>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
&nbsp;

Moving from the bottom to the top, this is how the flow looks:

1. At the bottom are Go packages in the gno codebase. A **gnoclient.Client** supports communication with the remote Gno.land node with methods like Call to call a realm function. The Gno codebase also has **keys.Keybase** to support a wallet stored on the local device with methods like CreateAccount.
2. These methods are called directly from the next level up by the **GnoMobile** Go code. A Go object can’t be passed through the gRPC interface, so the GnoMobile Go code maintains a persistent gnoclient.Client object, which is accessed by gRPC calls. The GnoMobile API functions are registered by an amino package.go file and the generated Protobuf files are used to configure the gRPC server.
3. Finally, at the top of the diagram, the **gRPC client in the mobile app** communicates with the GnoMobile gRPC server over a local connection using Protobuf messages. A gRPC call can either return an immediate result (for example, GetKeyCount) or an asynchronous gRPC stream object, which can return delayed results (for example, a Call to a remote realm function). The gRPC framework uses the Protobuf API to generate convenient API functions in the mobile app’s [preferred language](https://grpc.io/docs/languages) (React Native, Java, Swift, etc.).

# How GnoMobile benefits builders

The first version of the framework will include three main sets of features:

1. **Blockchain Operations**: These refer to the core block of functions that the apps need to interact with the blockchain. Things like the gnoclient API to effectively bring the benefits of the Gno framework on mobile, the gas estimation interface and calling realm functions, querying a blockchain node (and more) are included here.
2. **Wallet**: As the name suggests, here we have all the standard wallet operations like create or delete an account, set the recovery phrase, account balance, and so on.
3. **Toolkit**: We want to make it as easy as possible for devs to start building apps with our framework, so we’ll provide them with install instructions, example apps, and more technical stuff like genproto options to support gRPC and helper functions to parse the render output.

Those should be enough to allow builders to get started on using and experimenting with Gno mobile apps.

- *Support for secure p2p communication, even when the Internet is down?*
- *Yes, please!*

Something that is not necessarily essential for V1, but for sure will open the doors to some powerful capabilities later on is to add an interface and a constructor to adapt the communication transport. This will make it possible for devs to incorporate other tools like Wesh Network and give their apps the ability to securely and reliably send messages even in very poor network conditions. But that’s a story for another time.

# When will GnoMobile be ready?

V1 is planned for release in mid-December 2023.

Until then, you can check out our progress [here](https://github.com/gnolang/hackerspace/issues/28).

Got feedback or want to drop us a question? Ask away on our [repo](https://github.com/gnolang/gnomobile/issues).

# What does the future look like beyond V1?

We see a lot of potential directions for GnoMobile after the initial release that will improve the user experience, extend its functionality, and make GnoMobile even more secure. We’re still scratching the surface in terms of how far we can take its development, and we look forward to working on further iterations and improvements. Some of our ideas for the future beyond V1 include:

1. Making it easier for developers to **build** **desktop apps** **and** **browser extensions**:
2. Through GnoMobile, we can gradually enable “desktop” devs to use our React Native gRPC interface to write desktop applications while using existing functionality from the core Go code. This way, developers will not necessarily have to learn Go to leverage its advantages.
3. Browser extensions are usually written in JavaScript in the same way as in React Native. This opens the door to getting the benefits of Go via the GnoMobile framework. Otherwise, you’d have to either make the Go code run inside the browser extension (which is not easy) or use a remote server (which is not pretty).
4. Making it possible to **execute smart contracts directly from mobile**.

*Why is this important?*

If you want to add a new message to a blockchain, you need to actually interact with it (the blockchain) and update its state with the new message. However, if you just want to browse through the messages, you can execute the Render function locally without needing to use your network and, at the same time, get the results much faster. This is because the node runs locally on the mobile device without needing to spend crypto coins to get a remote node to do the operation for you.

Gno nodes run on GnoVMs (gnovm), and for the moment, these are only available on desktops. We believe it is possible to make them available on mobile as well, but we need to find clever ways to overcome the constraints of mobile devices (like putting the apps in the background (iOS), addressing network bandwidth limitations, and so on).

1. Developing a **decentralized push notification service** for *both* mobile and desktop apps. Getting notifications is now a standard (and very important) functionality of centralized apps. Technically, this happens via a central server. Naturally, having a centralized server is not possible for a p2p app, but there are other ways to implement notifications, and we are considering including them in the GnoMobile framework.
2. Making it possible for decentralized apps to **interact with the blockchain even if the network connection is poor or virtually unavailable**. Through the **[Wesh Network](https://wesh.network/)** protocol, we are opening up the possibility of using alternative transport mediums to exchange messages between peers in an asynchronous but reliable manner in off-grid environments. Enabling reliable, secure, and censorship-resistant communication is our main cause at Berty Technologies. We want to open the door for p2p users to send messages and interact even in extreme situations or adverse scenarios, and Wesh Network is built specifically for this purpose. It is only natural to make it easier for developers to use it through the GnoMobile framework.
3. Advancing **edge networking for enhanced blockchain resilience**. Edge networking refers to bringing functionality like computing power or storage closer to the user so that they don't need to travel through the whole Internet to interact with a server. The same edge concept can be applied to bring the necessary services to interact with the blockchain closer to each p2p user. For example, hosting a copy of the blockchain so a user can sync it or even execute smart contracts. Having these fundamental services closer to the p2p users is especially important in the case of mobile apps. We want to offer developers the possibility of taking advantage of the edge networking benefits by allowing them to use, for instance, network address redirections or special HTTP headers in the configuration of their applications.

In all honesty, it’s hard not to get excited about all the different possibilities that lie ahead for GnoMobile, but we’re keeping our focus on shipping V1 for now and collecting feedback from the community. After that, well, we hope you’ll stick around to see what happens next!
